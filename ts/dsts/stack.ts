export class Stack<T> {
    private readonly data: T[];
    private numberOfElements: number;

    constructor() {
        this.numberOfElements = 0;
        this.data = new Array<T>();
    }

    public get count() {
        return this.numberOfElements;
    }

    public Add(data: T): void {
        this.data.unshift(data);
        this.numberOfElements++;
    }

    public Remove(): T | null {
        const element = this.data.shift();
        if (element == undefined) {
            return null;
        }

        this.numberOfElements--;
        return element;
    }
}