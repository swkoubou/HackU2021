export interface User {
    id: string;
    displayName: string;
}

export class DummyUser implements User {
    id = "u_dummy";
    displayName = "Mr. Dummy";
}
