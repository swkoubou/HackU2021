import { User } from '../models/user_model'
import { UserRepository } from '../repositories/user_repository'

class UserRepositoryOffLineImpl implements UserRepository {
  private userDB: Map<string, User> = new Map()

  public Create(id: string, displayName: string) {
    if (this.userDB.has(id)) {
      throw new Error(`既に登録されているIDです。${id}`)
    }

    const user: User = {
      id: id,
      displayName: displayName,
    }

    return user
  }

  public Get(id: string) {
    if (this.userDB.has(id)) {
      return this.userDB.get(id)
    }

    throw new Error(`登録されていないIDです。: ${id}`)
  }

  public Update(id: string, displayName: string) {
    if (this.userDB.has(id)) {
      const newUserData: User = {
        id: id,
        displayName: displayName,
      }
      this.userDB.set(id, newUserData)
      return newUserData
    }

    throw new Error(`登録されていないIDです。: ${id}`)
  }

  public Delete(id: string) {
    return this.userDB.delete(id)
  }
}
