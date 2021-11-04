import { Question } from '../models/question'
import { User } from '../models/user'
import { QuestionRepository } from '../repositories/question_repository'

class OfflineQuestionRepositoryImpl implements QuestionRepository {
  public Create(
    author: User,
    title: string,
    tags: string[],
    type: Number,
    body: string,
    answers: string[],
    color: string
  ): Question {
    const question: Question = {
      id: 'questionId',
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

    return question
  }

  public Get(id: string): Question {
    const user: User = {
      displayName: 'displayName',
      id: 'userId',
    }

    const question: Question = {
      id: 'questionId',
      createdAt: new Date(),
      updatedAt: new Date(),
      author: user,
      title: 'title',
      tags: [],
      type: 0,
      body: 'body',
      answers: [],
      color: '#cee9e7',
    }
    return question
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
    const newQuestionData: Question = {
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

    return newQuestionData
  }

  public Delete(id: string): void {
    return
  }
}
