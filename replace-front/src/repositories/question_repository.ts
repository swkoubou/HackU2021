import { User } from '../models/user_model'
import { Question } from '../models/question_model'

export interface QuestionRepository {
  Create: (
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ) => Question
  Get: (id: string) => Question
  Update: (
    id: string,
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ) => Question
  Delete: (id: string) => void
}
