import { User } from './user'

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
