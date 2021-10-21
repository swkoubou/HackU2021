import {Question, DummyQuestion} from "../models/question";

export interface QuestionController {
    getLatestQuestion(): Question;
    getRandomQuestion(): Question;
    getQuestionWithTags(tags: string[]): Question;
}

export class OfflineQuestionController implements QuestionController {
    getLatestQuestion() {
        return new DummyQuestion();
    }
    getRandomQuestion() {
        return new DummyQuestion();
    }
    getQuestionWithTags(tags: string[]) {
        const dummy = new DummyQuestion();
        dummy.tags = tags;
        return dummy;
    }
}