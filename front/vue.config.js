if (process.env.API_KEY == null) {
  console.error('Error 環境変数「API_KEY」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_API_KEY = process.env.API_KEY
}

if (process.env.AUTH_DOMAIN == null) {
  console.error('Error 環境変数「AUTH_DOMAIN」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_AUTH_DOMAIN = process.env.AUTH_DOMAIN
}

if (process.env.PROJECT_ID == null) {
  console.error('Error 環境変数「PROJECT_ID」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_PROJECT_ID = process.env.PROJECT_ID
}

if (process.env.STORAGE_BUCKET == null) {
  console.error('Error 環境変数「STORAGE_BUCKET」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_STORAGE_BUCKET = process.env.STORAGE_BUCKET
}

if (process.env.MESSAGING_SENDER_ID == null) {
  console.error('Error 環境変数「MESSAGING_SENDER_ID」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_MESSAGING_SENDER_ID = process.env.MESSAGING_SENDER_ID
}

if (process.env.APP_ID == null) {
  console.error('Error 環境変数「APP_ID」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_APP_ID = process.env.APP_ID
}
