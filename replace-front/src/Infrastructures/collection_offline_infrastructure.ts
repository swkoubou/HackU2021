import { Collection } from '../models/collection_model'
import { Question } from '../models/question_model'
import { User } from '../models/user_model'
import { CollectionRepository } from '../repositories/collection_repository'

class CollectionRepositoryOfflineImpl implements CollectionRepository {
  private collectionDB: Map<string, Collection> = new Map()

  public Create(
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ) {
    if (this.collectionDB.has(id)) {
      throw new Error(`既に登録されているIDです。: ${id}`)
    }

    const collection: Collection = {
      id: id,
      createdAt: new Date(),
      updatedAt: new Date(),
      author: author,
      title: title,
      description: description,
      questions: questions,
      color: color,
    }

    this.collectionDB.set(id, collection)

    return collection
  }

  public Get(id: string) {
    if (this.collectionDB.has(id)) {
      return this.collectionDB.get(id)
    }

    throw new Error(`登録されていないIDです。: ${id}`)
  }

  public Update(
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ) {
    if (this.collectionDB.has(id)) {
      const oldCollectionData: Collection = this.collectionDB.get(id)!
      const createdAt: Date = oldCollectionData.createdAt

      const newCollectionData: Collection = {
        id: id,
        createdAt: createdAt,
        updatedAt: new Date(),
        author: author,
        title: title,
        description: description,
        questions: questions,
        color: color,
      }

      this.collectionDB.set(id, newCollectionData)

      return newCollectionData
    }

    throw new Error(`登録されていないIDです。: ${id}`)
  }

  public Delete(id: string) {
    return this.collectionDB.delete(id)
  }
}
