import { Question } from '../models/question'
import { QuestionRepository } from '../repositories/question_repository'

interface QuestionService {
  // IDを使って、問題を取得する
  GetQuestionByID: (id: string) => Question
  // 答え合わせをする
  CheckAnswers: (question: Question, answers: string[]) => boolean
}

export class QuestionServiceImpl implements QuestionService {
  private repository: QuestionRepository
  public constructor(repository: QuestionRepository) {
    this.repository = repository
  }

  GetQuestionByID(id: string): Question {
    return this.repository.Get(id)
  }

  CheckAnswers(question: Question, answers: string[]): boolean {
    if (question.answers.length != answers.length) {
      // 回答の数が一致しない
      return false
    }

    // １つ１つあっているか確認してく
    for (let i = 1; i < answers.length; i++) {
      if (question.answers[i] != answers[i]) {
        return false
      }
    }

    return true
  }
}
