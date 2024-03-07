export class LinkedNode<T> {
    public value: T;
    public next: LinkedNode<T> | null;

    constructor(value: T) {
        this.value = value;
        this.next = null
    }
}

export class LinkedList<T> {
    public head: LinkedNode<T> | null;
    public count: number;

    constructor() {
        this.head = null;
        this.count = 0;
    }

    public append(value: T) {
        const node = new LinkedNode<T>(value);

        if (this.head == null) {
            this.head = node;
            this.count++;
            return;
        }

        let current = this.head;

        while (current.next != null) {
            current = current.next;
        }

        current.next = node;
        this.count++;
    }

    public remove(value: T) {
        const head = this.head;
        if (head == null) {
            return;
        }

        if (head.value == value) {
            const value = head.value;
            this.head = head.next;
            this.count--;

            return value;
        }

        let previous = head;
        let current = head.next;

        if (current == null) {
            return;
        }

        while (current != null && current.next != null && current.value != value) {
            previous = current;
            current = current.next;
        }

        const foundValue = current.value;
        if (foundValue == value) {
            this.count--;
            previous.next = current.next;
            return value;
        }

        return;
    }

    public findFirst(value: T) {
        let current = this.head;
        while (current != null && current.value != value) {
            current = current.next;
        }

        return current != null && current.value || null;
    }
}