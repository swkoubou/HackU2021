import { Collection } from '../models/collection_model'
import { Question } from '../models/question_model'
import { User } from '../models/user_model'
import { CollectionRepository } from '../repositories/collection_repository'
import * as error_infrastructures from './error_infrastructures'

class OfflineCollectionRepository implements CollectionRepository {
  private collectionDB: Map<string, Collection> = new Map()
  private autoIncrementId: number = 0

  private newId(): string {
    this.autoIncrementId++
    return `cid_${this.autoIncrementId}`
  }

  public Create(
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ): Collection {
    const id = this.newId()
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

  public Get(id: string): Collection {
    if (this.collectionDB.has(id)) {
      return this.collectionDB.get(id)!
    }
    throw new error_infrastructures.OfflineDBCollectionNotFoundError(id)
  }

  public Update(
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ): Collection {
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

    throw new error_infrastructures.OfflineDBCollectionNotFoundError(id)
  }

  public Delete(id: string): void {
    if (this.collectionDB.has(id)) {
      this.collectionDB.delete(id)
      return
    }
    throw new error_infrastructures.OfflineDBCollectionNotFoundError(id)
  }
}
