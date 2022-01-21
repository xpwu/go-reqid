package reqid

import (
  "context"
  "crypto/md5"
  "encoding/hex"
  "github.com/google/uuid"
  "time"
)

type reqIDKey struct {}

const (
  HeaderKey = "X-Req-Id"
)


func WithCtx(parent context.Context) (ctx context.Context, id string) {
  switch value := parent.Value(reqIDKey{}).(type) {
  case string:
    return parent, value
  default:
    id = RandomID()
    ctx = context.WithValue(parent, reqIDKey{},id)
    return
  }
}

func NewContext(parent context.Context, id string) context.Context {
  if value,ok := parent.Value(reqIDKey{}).(string); ok && value == id {
    return parent
  }
  return context.WithValue(parent, reqIDKey{}, id)
}

func RandomID() string {
  // 一般不会出错
  id,err := uuid.NewRandom()
  if err != nil {
    res := make([]byte, md5.Size*2)
    m := md5.Sum([]byte(time.Now().String() + err.Error()))
    hex.Encode(res, m[:])
    return string(res)
  }

  return id.String()
}


