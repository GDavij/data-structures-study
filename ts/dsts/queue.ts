export class Queue<T> {
    private readonly data: T[];
    private numberOfElements: number;

    constructor() {
        this.numberOfElements = 0;
        this.data = new Array<T>();
    }

    public get count() {
        return this.numberOfElements;
    }

    public enqueue(value: T): void {
        this.numberOfElements++;
        this.data.unshift(value)
    }

    public dequeue(): T | null {
        const value = this.data.pop();
        if (value == undefined) {
            return null;
        }

        this.numberOfElements--;
        return value;
    }
}