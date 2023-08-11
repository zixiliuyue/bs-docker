// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bscp.io/pkg/dal/table"
)

func newTemplateRelease(db *gorm.DB, opts ...gen.DOOption) templateRelease {
	_templateRelease := templateRelease{}

	_templateRelease.templateReleaseDo.UseDB(db, opts...)
	_templateRelease.templateReleaseDo.UseModel(&table.TemplateRelease{})

	tableName := _templateRelease.templateReleaseDo.TableName()
	_templateRelease.ALL = field.NewAsterisk(tableName)
	_templateRelease.ID = field.NewUint32(tableName, "id")
	_templateRelease.ReleaseName = field.NewString(tableName, "release_name")
	_templateRelease.ReleaseMemo = field.NewString(tableName, "release_memo")
	_templateRelease.Name = field.NewString(tableName, "name")
	_templateRelease.Path = field.NewString(tableName, "path")
	_templateRelease.FileType = field.NewString(tableName, "file_type")
	_templateRelease.FileMode = field.NewString(tableName, "file_mode")
	_templateRelease.User = field.NewString(tableName, "user")
	_templateRelease.UserGroup = field.NewString(tableName, "user_group")
	_templateRelease.Privilege = field.NewString(tableName, "privilege")
	_templateRelease.Signature = field.NewString(tableName, "signature")
	_templateRelease.ByteSize = field.NewUint64(tableName, "byte_size")
	_templateRelease.BizID = field.NewUint32(tableName, "biz_id")
	_templateRelease.TemplateSpaceID = field.NewUint32(tableName, "template_space_id")
	_templateRelease.TemplateID = field.NewUint32(tableName, "template_id")
	_templateRelease.Creator = field.NewString(tableName, "creator")
	_templateRelease.CreatedAt = field.NewTime(tableName, "created_at")

	_templateRelease.fillFieldMap()

	return _templateRelease
}

type templateRelease struct {
	templateReleaseDo templateReleaseDo

	ALL             field.Asterisk
	ID              field.Uint32
	ReleaseName     field.String
	ReleaseMemo     field.String
	Name            field.String
	Path            field.String
	FileType        field.String
	FileMode        field.String
	User            field.String
	UserGroup       field.String
	Privilege       field.String
	Signature       field.String
	ByteSize        field.Uint64
	BizID           field.Uint32
	TemplateSpaceID field.Uint32
	TemplateID      field.Uint32
	Creator         field.String
	CreatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (t templateRelease) Table(newTableName string) *templateRelease {
	t.templateReleaseDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t templateRelease) As(alias string) *templateRelease {
	t.templateReleaseDo.DO = *(t.templateReleaseDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *templateRelease) updateTableName(table string) *templateRelease {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint32(table, "id")
	t.ReleaseName = field.NewString(table, "release_name")
	t.ReleaseMemo = field.NewString(table, "release_memo")
	t.Name = field.NewString(table, "name")
	t.Path = field.NewString(table, "path")
	t.FileType = field.NewString(table, "file_type")
	t.FileMode = field.NewString(table, "file_mode")
	t.User = field.NewString(table, "user")
	t.UserGroup = field.NewString(table, "user_group")
	t.Privilege = field.NewString(table, "privilege")
	t.Signature = field.NewString(table, "signature")
	t.ByteSize = field.NewUint64(table, "byte_size")
	t.BizID = field.NewUint32(table, "biz_id")
	t.TemplateSpaceID = field.NewUint32(table, "template_space_id")
	t.TemplateID = field.NewUint32(table, "template_id")
	t.Creator = field.NewString(table, "creator")
	t.CreatedAt = field.NewTime(table, "created_at")

	t.fillFieldMap()

	return t
}

func (t *templateRelease) WithContext(ctx context.Context) ITemplateReleaseDo {
	return t.templateReleaseDo.WithContext(ctx)
}

func (t templateRelease) TableName() string { return t.templateReleaseDo.TableName() }

func (t templateRelease) Alias() string { return t.templateReleaseDo.Alias() }

func (t *templateRelease) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *templateRelease) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 17)
	t.fieldMap["id"] = t.ID
	t.fieldMap["release_name"] = t.ReleaseName
	t.fieldMap["release_memo"] = t.ReleaseMemo
	t.fieldMap["name"] = t.Name
	t.fieldMap["path"] = t.Path
	t.fieldMap["file_type"] = t.FileType
	t.fieldMap["file_mode"] = t.FileMode
	t.fieldMap["user"] = t.User
	t.fieldMap["user_group"] = t.UserGroup
	t.fieldMap["privilege"] = t.Privilege
	t.fieldMap["signature"] = t.Signature
	t.fieldMap["byte_size"] = t.ByteSize
	t.fieldMap["biz_id"] = t.BizID
	t.fieldMap["template_space_id"] = t.TemplateSpaceID
	t.fieldMap["template_id"] = t.TemplateID
	t.fieldMap["creator"] = t.Creator
	t.fieldMap["created_at"] = t.CreatedAt
}

func (t templateRelease) clone(db *gorm.DB) templateRelease {
	t.templateReleaseDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t templateRelease) replaceDB(db *gorm.DB) templateRelease {
	t.templateReleaseDo.ReplaceDB(db)
	return t
}

type templateReleaseDo struct{ gen.DO }

type ITemplateReleaseDo interface {
	gen.SubQuery
	Debug() ITemplateReleaseDo
	WithContext(ctx context.Context) ITemplateReleaseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITemplateReleaseDo
	WriteDB() ITemplateReleaseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITemplateReleaseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITemplateReleaseDo
	Not(conds ...gen.Condition) ITemplateReleaseDo
	Or(conds ...gen.Condition) ITemplateReleaseDo
	Select(conds ...field.Expr) ITemplateReleaseDo
	Where(conds ...gen.Condition) ITemplateReleaseDo
	Order(conds ...field.Expr) ITemplateReleaseDo
	Distinct(cols ...field.Expr) ITemplateReleaseDo
	Omit(cols ...field.Expr) ITemplateReleaseDo
	Join(table schema.Tabler, on ...field.Expr) ITemplateReleaseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITemplateReleaseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITemplateReleaseDo
	Group(cols ...field.Expr) ITemplateReleaseDo
	Having(conds ...gen.Condition) ITemplateReleaseDo
	Limit(limit int) ITemplateReleaseDo
	Offset(offset int) ITemplateReleaseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITemplateReleaseDo
	Unscoped() ITemplateReleaseDo
	Create(values ...*table.TemplateRelease) error
	CreateInBatches(values []*table.TemplateRelease, batchSize int) error
	Save(values ...*table.TemplateRelease) error
	First() (*table.TemplateRelease, error)
	Take() (*table.TemplateRelease, error)
	Last() (*table.TemplateRelease, error)
	Find() ([]*table.TemplateRelease, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.TemplateRelease, err error)
	FindInBatches(result *[]*table.TemplateRelease, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.TemplateRelease) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITemplateReleaseDo
	Assign(attrs ...field.AssignExpr) ITemplateReleaseDo
	Joins(fields ...field.RelationField) ITemplateReleaseDo
	Preload(fields ...field.RelationField) ITemplateReleaseDo
	FirstOrInit() (*table.TemplateRelease, error)
	FirstOrCreate() (*table.TemplateRelease, error)
	FindByPage(offset int, limit int) (result []*table.TemplateRelease, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITemplateReleaseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t templateReleaseDo) Debug() ITemplateReleaseDo {
	return t.withDO(t.DO.Debug())
}

func (t templateReleaseDo) WithContext(ctx context.Context) ITemplateReleaseDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t templateReleaseDo) ReadDB() ITemplateReleaseDo {
	return t.Clauses(dbresolver.Read)
}

func (t templateReleaseDo) WriteDB() ITemplateReleaseDo {
	return t.Clauses(dbresolver.Write)
}

func (t templateReleaseDo) Session(config *gorm.Session) ITemplateReleaseDo {
	return t.withDO(t.DO.Session(config))
}

func (t templateReleaseDo) Clauses(conds ...clause.Expression) ITemplateReleaseDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t templateReleaseDo) Returning(value interface{}, columns ...string) ITemplateReleaseDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t templateReleaseDo) Not(conds ...gen.Condition) ITemplateReleaseDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t templateReleaseDo) Or(conds ...gen.Condition) ITemplateReleaseDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t templateReleaseDo) Select(conds ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t templateReleaseDo) Where(conds ...gen.Condition) ITemplateReleaseDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t templateReleaseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITemplateReleaseDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t templateReleaseDo) Order(conds ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t templateReleaseDo) Distinct(cols ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t templateReleaseDo) Omit(cols ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t templateReleaseDo) Join(table schema.Tabler, on ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t templateReleaseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t templateReleaseDo) RightJoin(table schema.Tabler, on ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t templateReleaseDo) Group(cols ...field.Expr) ITemplateReleaseDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t templateReleaseDo) Having(conds ...gen.Condition) ITemplateReleaseDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t templateReleaseDo) Limit(limit int) ITemplateReleaseDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t templateReleaseDo) Offset(offset int) ITemplateReleaseDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t templateReleaseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITemplateReleaseDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t templateReleaseDo) Unscoped() ITemplateReleaseDo {
	return t.withDO(t.DO.Unscoped())
}

func (t templateReleaseDo) Create(values ...*table.TemplateRelease) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t templateReleaseDo) CreateInBatches(values []*table.TemplateRelease, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t templateReleaseDo) Save(values ...*table.TemplateRelease) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t templateReleaseDo) First() (*table.TemplateRelease, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.TemplateRelease), nil
	}
}

func (t templateReleaseDo) Take() (*table.TemplateRelease, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.TemplateRelease), nil
	}
}

func (t templateReleaseDo) Last() (*table.TemplateRelease, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.TemplateRelease), nil
	}
}

func (t templateReleaseDo) Find() ([]*table.TemplateRelease, error) {
	result, err := t.DO.Find()
	return result.([]*table.TemplateRelease), err
}

func (t templateReleaseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.TemplateRelease, err error) {
	buf := make([]*table.TemplateRelease, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t templateReleaseDo) FindInBatches(result *[]*table.TemplateRelease, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t templateReleaseDo) Attrs(attrs ...field.AssignExpr) ITemplateReleaseDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t templateReleaseDo) Assign(attrs ...field.AssignExpr) ITemplateReleaseDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t templateReleaseDo) Joins(fields ...field.RelationField) ITemplateReleaseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t templateReleaseDo) Preload(fields ...field.RelationField) ITemplateReleaseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t templateReleaseDo) FirstOrInit() (*table.TemplateRelease, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.TemplateRelease), nil
	}
}

func (t templateReleaseDo) FirstOrCreate() (*table.TemplateRelease, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.TemplateRelease), nil
	}
}

func (t templateReleaseDo) FindByPage(offset int, limit int) (result []*table.TemplateRelease, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t templateReleaseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t templateReleaseDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t templateReleaseDo) Delete(models ...*table.TemplateRelease) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *templateReleaseDo) withDO(do gen.Dao) *templateReleaseDo {
	t.DO = *do.(*gen.DO)
	return t
}