import { User } from '../models/user_model'
import { Question } from '../models/question_model'
import { Collection } from '../models/collection_model'

export interface CollectionRepository {
  Create: (
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ) => Collection | undefined
  Get: (id: string) => Collection | undefined
  Update: (
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ) => Collection
  Delete: (id: string) => boolean
}
