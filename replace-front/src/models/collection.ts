interface Collection {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    author: User;

    title: string;
    description: string;
    questions: Question[];
    color: string;
}