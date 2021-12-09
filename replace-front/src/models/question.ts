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
  answers: any
  color: string
}

export interface AnswerableQuestion extends Question {
  userAnswers: any
  CheckAnswers: () => boolean
}

export class FillInTheHolesQuestion implements AnswerableQuestion {
  id: string
  createdAt: Date
  updatedAt: Date
  author: User

  title: string
  tags: string[]
  type: Number
  body: string
  answers: any
  color: string
  userAnswers: string[]
  constructor(q: Question) {
    this.id = q.id
    this.createdAt = q.createdAt
    this.updatedAt = q.updatedAt
    this.author = q.author

    this.title = q.title
    this.tags = q.tags
    this.type = q.type
    this.body = q.body
    this.answers = q.answers as string[]
    this.color = q.color
    this.userAnswers = [] as string[]
  }

  CheckAnswers() {
    if (this.answers.length != this.userAnswers.length) {
      throw Error()
    }
    for (var i = 0; i < this.answers.length; i++) {
      if (this.answers[i] != this.userAnswers[i]) {
        return false
      }
    }
    return true
  }
}
