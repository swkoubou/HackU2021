import {User, DummyUser} from "../models/user";


export interface UserController {
    getUserWithId(id: string): User;
}

export class OfflineUserController implements UserController {
    getUserWithId(id: string): User {
        const dummy = new DummyUser();
        dummy.id = id;
        return dummy;
    }
}