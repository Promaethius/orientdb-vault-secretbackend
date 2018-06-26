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

// REF: KeepAlive pool.
tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: false,
}
client := &http.Client{Transport: tr}

type OrientDB struct {
  *orientdbConnectionProducer
  credsutil.CredentialsProducer
}

