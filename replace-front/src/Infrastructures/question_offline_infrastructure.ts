import { Question } from '../models/question_model'
import { User } from '../models/user_model'
import { QuestionRepository } from '../repositories/question_repository'

class OfflineQuestionRepository implements QuestionRepository {
  private questionDB: Map<string, Question> = new Map()
  private autoIncrementId: number = 0

  private newId(): string {
    this.autoIncrementId++
    return `qid_${this.autoIncrementId}`
  }

  public Create(
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ): Question {
    const id = this.newId()
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

  public Get(id: string): Question {
    if (this.questionDB.has(id)) {
      return this.questionDB.get(id)!
    }

    throw new OfflineDBQuestionNotFoundError(id)
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
  ): Question {
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

    throw new OfflineDBQuestionNotFoundError(id)
  }

  public Delete(id: string): void {
    if (this.questionDB.has(id)) {
      this.questionDB.delete(id)
      return
    }
    throw new OfflineDBQuestionNotFoundError(id)
  }
}
