import { User } from '../models/user'

export interface UserRepository {
  Create: (displayName: string) => User
  Get: (id: string) => User
  Update: (id: string, displayName: string) => User
  Delete: (id: string) => void
}
