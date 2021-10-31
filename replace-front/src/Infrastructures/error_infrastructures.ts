// 参考: http://exhikkii.com/2020/11/01/typescript-カスタムエラーの作成の仕方/

// 元になるエラー
class OfflineDBError implements Error {
  public name = 'OfflineDBError'
  constructor(public message: string) {
    if (typeof console !== 'undefined') {
      console.log(`name: ${this.name}, message: ${this.message}`)
    }
  }
  toString() {
    return `${this.name} ${this.message}`
  }
}

// ユーザーが見つからない
class OfflineDBUserNotFoundError extends OfflineDBError {
  public name = 'OfflineDBUserNotFoundError'
  constructor(public id: string) {
    super(`userId not found: ${id}`)
  }
}

// クエスチョンが見つからない
class OfflineDBQuestionNotFoundError extends OfflineDBError {
  public name = 'OfflineDBQuestionNotFoundError'
  constructor(public id: string) {
    super(`questionId not found: ${id}`)
  }
}

// コレクションが見つからない
class OfflineDBCollectionNotFoundError extends OfflineDBError {
  public name = 'OfflineDBCollectionNotFoundError'
  constructor(public id: string) {
    super(`collectionId not found: ${id}`)
  }
}
