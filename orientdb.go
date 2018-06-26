package orientdb

import (
	"context"
	"net/http"
	"time"
	"strings"
	"error"
	"encoding/json"
	"database/sql"

	// REF: Run() api.TLSconfig
	"github.com/hashicorp/vault/api"

	// REF: UsernameConfig
	"github.com/hashicorp/vault/builtin/logical/database/dbplugin"

	// REF: CredsUtil
	"github.com/hashicorp/vault/plugins/helper/database/credsutil"
)

// REF: Called when plugin is started
func Run(apiTLSConfig *api.TLSConfig) error {
	// Create keep-alive pool when plugin starts.
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
		
	dbType, err := // TODO: init OrientDB struct here, pass tr into structure for connection pool
	if err != nil {
		return err
	}

	plugins.Serve(dbType.(dbplugin.Database), apiTLSConfig)

	return nil
}

// REF: structure passed as plugin interface to vault.
type OrientDB struct {
	*orientdbConnectionProducer
	credsutil.CredentialsProducer
	
	// Init a client that uses the previous connection pool
	orientdbClient := &http.Client{Transport: tr}
	
	// REF: Init(ctx context.Context, config map[string]interface{}, verifyConnection bool) (saveConfig map[string]interface{}, err error)
	// REF: CreateUser(ctx context.Context, statements dbplugin.Statements, usernameConfig dbplugin.UsernameConfig, expiration time.Time) (username string, password string, err error)
	// REF: RenewUser(ctx context.Context, statements dbplugin.Statements, username string, expiration time.Time) error
	// REF: RevokeUser(ctx context.Context, statements dbplugin.Statements, username string) error
	// REF: RotateRootCredentials(ctx context.Context, statements []string) (map[string]interface{}, error)
	// REF: Close() error
}

func (c *OrientDB) Init(ctx context.Context, config map[string]interface{}, verifyConnection bool) (saveConfig map[string]interface{}, err error) {
	
}

// Type returns the name of this plugin
func (c *OrientDB) Type() (string, error) {
	return "OrientDB", nil
}

func (c *OrientDB) CreateUser(ctx context.Context, statements dbplugin.Statements, usernameConfig dbplugin.UsernameConfig, expiration time.Time) (username string, password string, err error) {
	c.Lock()
	defer c.Unlock()
}


func (c *OrientDB) RenewUser(ctx context.Context, statements dbplugin.Statements, username string, expiration time.Time) error {
	
}

func (c *OrientDB) RevokeUser(ctx context.Context, statements dbplugin.Statements, username string) error {
	
}

func (c *OrientDB) RotateRootCredentials(ctx context.Context, statements []string) (map[string]interface{}, error) {
	
}

func (c * OrientDB) Close() error {
	
}

