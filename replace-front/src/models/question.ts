import { User, DummyUser } from "./user";

export interface Question {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    author: User;

    title: string;
    tags: string[];
    type: Number;
    body: string;
    answers: string[];
    color: string;
}

export class DummyQuestion implements Question {
    id = "q_dummy";
    createdAt = new Date();
    updatedAt = new Date();
    author = new DummyUser();

    title = "Q. Dummy"
    tags = ["tag01", "tag02", "tag03"];
    type = 0;
    body = "du, du, dummy";
    answers = ["ans01", "ans02", "ans03"];
    color = "#6897BB"; // blue
}