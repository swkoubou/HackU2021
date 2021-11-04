import { User } from '../models/user'
import { UserRepository } from '../repositories/user'
import * as error_infrastructures from '../errors/infrastructures'

class OfflineUserRepository implements UserRepository {
  private userDB: Map<string, User> = new Map()
  private autoIncrementId: number = 0

  private newId(): string {
    this.autoIncrementId++
    return `uid_${this.autoIncrementId}`
  }

  public Create(displayName: string): User {
    const id = this.newId()
    const user: User = {
      id: id,
      displayName: displayName,
    }

    return user
  }

  public Get(id: string): User {
    if (!this.userDB.has(id)) {
      throw new error_infrastructures.OfflineDBUserNotFoundError(id)
    }

    return this.userDB.get(id)!
  }

  public Update(id: string, displayName: string): User {
    if (!this.userDB.has(id)) {
      throw new error_infrastructures.OfflineDBUserNotFoundError(id)
    }

    const newUserData: User = {
      id: id,
      displayName: displayName,
    }
    this.userDB.set(id, newUserData)
    return newUserData
  }

  public Delete(id: string): void {
    if (!this.userDB.has(id)) {
      throw new error_infrastructures.OfflineDBUserNotFoundError(id)
    }
    this.userDB.delete(id)
    return
  }
}
