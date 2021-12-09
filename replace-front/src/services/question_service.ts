import { Question } from '../models/question'
import { QuestionRepository } from '../repositories/question_repository'
import { NotAllowedCharacterError } from '../errors/service_error'

interface QuestionService {
  // IDを使って、問題を取得する
  GetQuestionByID: (id: string) => Question
  // 最新の問題をn個取得する
  // GetLatestQuestion: (n: Number) => Question[]
}

export class QuestionServiceImpl implements QuestionService {
  private repository: QuestionRepository
  public constructor(repository: QuestionRepository) {
    this.repository = repository
  }

  // アルファベット、数字、表示可能な記号かどうか
  private isPrintableAscii(text: string): boolean {
    const printableAsciiRegex = /^[\x20-\x7F]*$/
    return printableAsciiRegex.test(text)
  }

  GetQuestionByID(id: string): Question {
    if (!this.isPrintableAscii(id)) {
      // 不正な文字列が検知された
      throw new NotAllowedCharacterError(id)
    }
    return this.repository.Get(id)
  }
}
