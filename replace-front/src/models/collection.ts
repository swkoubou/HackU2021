import {User, DummyUser} from "./user";
import { Question, DummyQuestion } from "./question";

export interface Collection {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    author: User;

    title: string;
    description: string;
    questions: Question[];
    color: string;
}

export class DummyCollection implements Collection {
    id = "c_dummy";
    createdAt = new Date();
    updatedAt = new Date();
    author = new DummyUser();

    title = "D-Collection"
    description = "this is dummy."
    questions = [new DummyQuestion()];
    color = "#f9cb9c"; // orange
}