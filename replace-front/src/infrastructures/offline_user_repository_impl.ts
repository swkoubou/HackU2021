import { User } from '../models/user'
import { UserRepository } from '../repositories/user_repository'

export function CreateUserRepository(): UserRepository {
  return new OfflineUserRepositoryImpl()
}

class OfflineUserRepositoryImpl implements UserRepository {
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
