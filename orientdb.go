package orientdb

import (
	"context"
	"net/http"
	"time"
	"strings"
	"error"

	// REF: Run() api.TLSconfig
	"github.com/hashicorp/vault/api"

	// REF: dbplugin.UsernameConfig https://github.com/hashicorp/vault/blob/master/builtin/logical/database/dbplugin/database.proto
	"github.com/hashicorp/vault/builtin/logical/database/dbplugin"
	/*
	message UsernameConfig {
		string DisplayName = 1;
		string RoleName = 2;
	}
	*/

	// REF: credsutil.CredentialsProducer https://github.com/hashicorp/vault/blob/master/plugins/helper/database/credsutil/credsutil.go
	"github.com/hashicorp/vault/plugins/helper/database/credsutil"
	/*
	type CredentialsProducer interface {
		GenerateUsername(usernameConfig dbplugin.UsernameConfig) (string, error)
		GeneratePassword() (string, error)
		GenerateExpiration(ttl time.Time) (string, error)
	}
	*/
)

// REF: Called when plugin is started by ./orientdb-vault-secretbackend/main.go
func Run(apiTLSConfig *api.TLSConfig) error {
	// Create keep-alive pool when plugin starts.
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
		
	dbType, err := // TODO(1): init OrientDB struct here, pass tr into structure for connection pool
	if err != nil {
		return err
	}

	plugins.Serve(dbType.(dbplugin.Database), apiTLSConfig)

	return nil
}

// REF: structure passed as plugin interface to vault.
type OrientDB struct {
	credsutil.CredentialsProducer
	
	// Init a client that uses the previous connection pool
	orientdbClient
	
	// REF: Init(ctx context.Context, config map[string]interface{}, verifyConnection bool) (saveConfig map[string]interface{}, err error)
	// REF: CreateUser(ctx context.Context, statements dbplugin.Statements, usernameConfig dbplugin.UsernameConfig, expiration time.Time) (username string, password string, err error)
	// REF: RenewUser(ctx context.Context, statements dbplugin.Statements, username string, expiration time.Time) error
	// REF: RevokeUser(ctx context.Context, statements dbplugin.Statements, username string) error
	// REF: RotateRootCredentials(ctx context.Context, statements []string) (map[string]interface{}, error)
	// REF: Close() error
}

func (c *OrientDB) Init(ctx context.Context, config map[string]interface{}, verifyConnection bool) (saveConfig map[string]interface{}, err error) {
	// TODO(1): tr will come from type init called from Run()
	orientdbClient := &http.Client{Transport: tr}
}

// Type returns the name of this plugin in a string
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
	// No idea if this is useful. Much easier to just close the connection body after a successful transaction.
	return nil
}

