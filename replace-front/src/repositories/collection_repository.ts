import { User } from '../models/user_model'
import { Question } from '../models/question_model'
import { Collection } from '../models/collection_model'

export interface CollectionRepository {
  Create: (
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ) => Collection
  Get: (id: string) => Collection
  Update: (
    id: string,
    author: User,
    title: string,
    description: string,
    questions: Question[],
    color: string
  ) => Collection
  Delete: (id: string) => void
}
