package repositories

import (
	"context"
	"reflect"

	"githubcom/kosatnkn/web-page-analyzer-api/app/adapters"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/boundary/repositories"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
	"githubcom/kosatnkn/web-page-analyzer-api/externals/repositories/errors"
	"github.com/kosatnkn/req"
	"github.com/kosatnkn/req/filter"
	"github.com/kosatnkn/req/paginator"
)

// SampleMySQLRepository is an example repository that implements test database functionality.
type SampleMySQLRepository struct {
	db adapters.DBAdapterInterface
	*filter.FilterRepositoryFacilitator
	*paginator.PaginatorRepositoryFacilitator
}

// NewSampleMySQLRepository creates a new instance of the repository.
func NewSampleMySQLRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {
	// Filter mappings are used to establish the connection between the filter and the underlying
	// table field. In addition to that an operator can be optionally assigned for the filter.
	// When an operator is not specified the mapping defaults to `equal` operator.
	//
	// In order for a filter to be assigned the value of `filter.Filter.Name` should match with
	// one of the keys in the `filterMap`.
	filterMap := map[string][]string{
		"NameContain": {"name", req.SelectLike},
	}

	return &SampleMySQLRepository{
		db:                             dbAdapter,
		FilterRepositoryFacilitator:    filter.NewFilterRepositoryFacilitator(req.DBMySQL, filterMap),
		PaginatorRepositoryFacilitator: paginator.NewPaginatorRepositoryFacilitator(req.DBMySQL),
	}
}

// Get retrieves a collection of Samples.
func (repo *SampleMySQLRepository) Get(ctx context.Context, fts []filter.Filter, pgn paginator.Paginator) ([]entities.Sample, error) {
	query := `SELECT id, name, password
				FROM sample`

	// add filters to query
	query, params, err := repo.WithFilters(query, fts)
	if err != nil {
		return nil, errors.ErrQuery(err)
	}

	// add pagination
	query = repo.WithPagination(query, pgn)

	result, err := repo.db.Query(ctx, query, params)
	if err != nil {
		return nil, errors.ErrQuery(err)
	}

	return repo.mapResult(result)
}

// GetByID retrieves a single Sample.
func (repo *SampleMySQLRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {
	// DBAdapters in Catalyst supports named parameters and you don't have to
	// worry about the order in which those parameters are declared in the
	// query and in the parameters map. The DBAdapter will take care of that.
	query := `SELECT id, name, password
				FROM sample
				WHERE id=?id`

	parameters := map[string]interface{}{
		"id": id,
	}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return entities.Sample{}, errors.ErrQuery(err)
	}

	mapped, err := repo.mapResult(result)
	if err != nil {
		return entities.Sample{}, errors.ErrQuery(err)
	}
	if len(mapped) == 0 {
		return entities.Sample{}, nil
	}

	return mapped[0], nil
}

// Add adds a new sample.
func (repo *SampleMySQLRepository) Add(ctx context.Context, sample entities.Sample) error {
	query := `INSERT INTO sample
				(name, password)
				VALUES(?name, ?password)`

	parameters := map[string]interface{}{
		"name":     sample.Name,
		"password": sample.Password,
	}

	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return errors.ErrQuery(err)
	}

	return nil
}

// Edit updates an existing sample identified by the id.
func (repo *SampleMySQLRepository) Edit(ctx context.Context, sample entities.Sample) error {
	query := `UPDATE sample
				SET name=?name, password=?password
				WHERE id=?id`

	parameters := map[string]interface{}{
		"id":       sample.ID,
		"name":     sample.Name,
		"password": sample.Password,
	}

	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return errors.ErrQuery(err)
	}

	return nil
}

// Delete deletes an existing sample identified by id.
func (repo *SampleMySQLRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM sample
				WHERE id=?id`

	parameters := map[string]interface{}{
		"id": id,
	}

	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return errors.ErrQuery(err)
	}

	return nil
}

// mapResult maps the result to entities.
func (repo *SampleMySQLRepository) mapResult(result []map[string]interface{}) (samples []entities.Sample, err error) {
	// Applying type assertion in this manner will result in a panic when the db data structure changes.
	// This defer recover pattern is used to recover from the panic and to return an error instead.
	// Notice the use of `named returned values` for this function (without which the recover pattern will not work).
	defer func() {
		if r := recover(); r != nil {
			err = errors.ErrQuery(r.(error))
		}
	}()

	for _, row := range result {
		samples = append(samples, entities.Sample{
			ID:   uint64(row["id"].(int64)),
			Name: string(row["name"].([]byte)),
		})
	}

	return samples, nil
}

// fallbackToNil returns nil if the value provided is the zero value of that type.
//
// Using 'nil' instead of zero values of types ensures that a 'NULL' is inserted to the db field.
// This is helpful when a field of a table is set to have a 'NULL' value when a value is not assigned to it.
func (repo *SampleMySQLRepository) fallbackToNil(val interface{}) interface{} {
	v := reflect.ValueOf(val)
	if v.IsZero() {
		return nil
	}

	return val
}

// convertToBool convert integer values to boolean.
func (repo *SampleMySQLRepository) convertToBool(val int64) bool {
	return val != 0
}

// getInsertID returns the id for the inserted record.
func (repo *SampleMySQLRepository) getInsertID(data []map[string]interface{}) int64 {

	if len(data) == 0 {
		return 0
	}

	return data[0]["last_insert_id"].(int64)
}

// getInsertID returns the id for the inserted record.
func (repo *SampleMySQLRepository) getAffectedRows(data []map[string]interface{}) int64 {
	if len(data) == 0 {
		return 0
	}

	return data[0]["affected_rows"].(int64)
}
