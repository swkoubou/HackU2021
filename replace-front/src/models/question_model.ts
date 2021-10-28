import { User } from './user_model'

export interface Question {
  id: string
  createdAt: Date
  updatedAt: Date
  author: User

  title: string
  tags: string[]
  type: Number
  body: string
  answers: string[]
  color: string
}
