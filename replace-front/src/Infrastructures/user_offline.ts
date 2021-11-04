import { User } from '../models/user'
import { UserRepository } from '../repositories/user'

class OfflineUserRepository implements UserRepository {
  public Create(displayName: string): User {
    const user: User = {
      id: 'userId',
      displayName: displayName,
    }

    return user
  }

  public Get(id: string): User {
    const user: User = {
      displayName: 'displayName',
      id: 'userId',
    }

    return user
  }

  public Update(id: string, displayName: string): User {
    const newUserData: User = {
      id: id,
      displayName: displayName,
    }

    return newUserData
  }

  public Delete(id: string): void {
    return
  }
}
