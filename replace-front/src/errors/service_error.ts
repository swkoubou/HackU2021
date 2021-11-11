export abstract class ServiceError extends Error {
  public abstract override readonly name: string
  constructor(message: string) {
    super(message)
  }
}

export class NotAllowedCharacterError extends ServiceError {
  public name: string = 'NotAllowedCharacterError'
  public invalidText: string
  constructor(invalidText: string) {
    super(
      `This string contains characters that are not allowed: ${invalidText}`
    )
    this.invalidText = invalidText
  }
}
