import { User } from '../models/user'
import { Question } from '../models/question'
import { Collection } from '../models/collection'

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
