import { Collection, DummyCollection } from "../models/collection";

export interface CollectionController {
    getCollection(): Collection;
}

export class OfflineCollectionController implements CollectionController {
    getCollection(): Collection {
        return new DummyCollection();
    }
}