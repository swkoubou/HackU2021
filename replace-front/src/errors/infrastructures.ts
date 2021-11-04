// 参考:
// http://exhikkii.com/2020/11/01/typescript-カスタムエラーの作成の仕方/
// https://stackoverflow.com/questions/41102060/typescript-extending-error-class
// https://qiita.com/suema0331/items/374c0757aa00b37d98bd

// 元になるエラー
export abstract class OfflineDBError extends Error {
  public abstract override readonly name: string
  constructor(message: string) {
    super(message)
  }
}

// ユーザーが見つからない
export class OfflineDBUserNotFoundError extends OfflineDBError {
  public name = 'OfflineDBUserNotFoundError'
  constructor(id: string) {
    super(`userId not found: ${id}`)
  }
}

// クエスチョンが見つからない
export class OfflineDBQuestionNotFoundError extends OfflineDBError {
  public name = 'OfflineDBQuestionNotFoundError'
  constructor(id: string) {
    super(`questionId not found: ${id}`)
  }
}

// コレクションが見つからない
export class OfflineDBCollectionNotFoundError extends OfflineDBError {
  public name = 'OfflineDBCollectionNotFoundError'
  constructor(id: string) {
    super(`collectionId not found: ${id}`)
  }
}
