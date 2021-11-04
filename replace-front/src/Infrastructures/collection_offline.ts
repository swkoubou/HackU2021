import { Collection } from '../models/collection'
import { Question } from '../models/question'
import { User } from '../models/user'
import { CollectionRepository } from '../repositories/collection'
import { OfflineDBCollectionNotFoundError } from '../errors/infrastructures'

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
    if (!this.collectionDB.has(id)) {
      throw new OfflineDBCollectionNotFoundError(id)
    }
    return this.collectionDB.get(id)!
  }

  public Update(
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ): Collection {
    if (!this.collectionDB.has(id)) {
      throw new OfflineDBCollectionNotFoundError(id)
    }

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

  public Delete(id: string): void {
    if (!this.collectionDB.has(id)) {
      throw new OfflineDBCollectionNotFoundError(id)
    }
    this.collectionDB.delete(id)
    return
  }
}
