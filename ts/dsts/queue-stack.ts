class QueueStackNode<T> {
    public value: T
    public next: QueueStackNode<T> | null

    constructor(value: T) {
        this.value = value;
        this.next = null;
    }
}

export class QueueStack<T> {
    public head: QueueStackNode<T> | null;

    constructor() {
        this.head = null;
    }

    public unShift(value: T): void {
        const node = new QueueStackNode<T>(value);
        node.next = this.head;
        this.head = node;
    }

    public push(value: T): void {
        const node = new QueueStackNode<T>(value);
        if (this.head == null) {
            this.head = node;
            return;
        }

        let current = this.head;
        while (current.next != null) {
            current = current.next;
        }

        current.next = node;
        return;
    }

    public shift(): T | null {
        if (this.head == null) {
            return null;
        }

        let result = this.head;
        this.head = result.next;
        return result.value;
    }

    public pop(): T | null {
        if (this.head == null) {
            return null;
        }

        let previous = this.head;
        let current = previous.next;

        if (current == null) {
            const result = previous.value;
            this.head = null;
            return result;
        }

        while (current.next != null) {
            previous = current;
            current = current!.next;
        }

        const result = previous.next!.value;
        previous.next = null;
        return result;

    }

}

