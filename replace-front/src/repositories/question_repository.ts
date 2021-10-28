import { User } from '../models/user_model'
import { Question } from '../models/question_model'

export interface QuestionRepository {
  Create: (
    id: string,
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ) => Question | undefined
  Get: (id: string) => Question | undefined
  Update: (
    id: string,
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ) => Question | undefined
  Delete: (id: string) => boolean
}
