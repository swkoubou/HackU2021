import { Collection } from '../models/collection'
import { Question } from '../models/question'
import { User } from '../models/user'
import { CollectionRepository } from '../repositories/collection_repository'

export function CreatCollectionRepository(): CollectionRepository {
  return new OfflineCollectionRepositoryImpl()
}

class OfflineCollectionRepositoryImpl implements CollectionRepository {
  public Create(
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ): Collection {
    const collection: Collection = {
      id: 'collectionId',
      createdAt: new Date(),
      updatedAt: new Date(),
      author: author,
      title: title,
      description: description,
      questions: questions,
      color: color,
    }
    return collection
  }

  public Get(id: string): Collection {
    const user: User = {
      displayName: 'displayName',
      id: 'userId',
    }

    const question: Question = {
      id: 'questionId',
      createdAt: new Date(),
      updatedAt: new Date(),
      author: user,
      title: 'title',
      tags: [],
      type: 0,
      body: 'body',
      answers: [],
      color: '#cee9e7',
    }

    const collection: Collection = {
      id: id,
      createdAt: new Date(),
      updatedAt: new Date(),
      author: user,
      title: 'title',
      description: '',
      questions: [question],
      color: '#cee9e7',
    }

    return collection
  }

  public Update(
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ): Collection {
    const newCollectionData: Collection = {
      id: id,
      createdAt: new Date(),
      updatedAt: new Date(),
      author: author,
      title: title,
      description: description,
      questions: questions,
      color: color,
    }

    return newCollectionData
  }

  public Delete(id: string): void {
    return
  }
}
