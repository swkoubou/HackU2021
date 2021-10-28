import { Question } from '../models/question_model'
import { User } from '../models/user_model'
import { QuestionRepository } from '../repositories/question_repository'

class QuestionRepositoryOfflineImpl implements QuestionRepository {
  private questionDB: Map<string, Question> = new Map()

  public Create(
    id: string,
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ) {
    if (this.questionDB.has(id)) {
      throw new Error(`既に登録されているIDです。: ${id}`)
    }

    const question: Question = {
      id: id,
      createdAt: new Date(),
      updatedAt: new Date(),
      author: author,
      title: title,
      tags: tags,
      type: type,
      body: body,
      answers: answers,
      color: color,
    }

    this.questionDB.set(id, question)

    return question
  }

  public Get(id: string) {
    if (this.questionDB.has(id)) {
      return this.questionDB.get(id)
    }

    throw new Error(`登録されていないIDです。: ${id}`)
  }

  public Update(
    id: string,
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ) {
    if (this.questionDB.has(id)) {
      const oldQuestionData: Question = this.questionDB.get(id)!
      const createdAt: Date = oldQuestionData.createdAt

      const newQuestionData: Question = {
        id: id,
        createdAt: createdAt,
        updatedAt: new Date(),
        author: author,
        title: title,
        tags: tags,
        type: type,
        body: body,
        answers: answers,
        color: color,
      }

      this.questionDB.set(id, newQuestionData)

      return newQuestionData
    }

    throw new Error(`登録されていないIDです。: ${id}`)
  }

  public Delete(id: string) {
    return this.questionDB.delete(id)
  }
}
