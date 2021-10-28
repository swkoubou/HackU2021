import { User } from '../models/user_model'

export interface UserRepository {
  Create: (id: string, displayName: string) => User | undefined
  Get: (id: string) => User | undefined
  Update: (id: string, displayName: string) => User | undefined
  Delete: (id: string) => boolean
}
