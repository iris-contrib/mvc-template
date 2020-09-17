export interface Counter {
    value: number
}

export function getCounter(): Promise<Counter> {
    return fetch('/api/counter')
        .then(res => res.json())
        .then(res => {
            return res as Counter;
        })
}

export function incrementCounter(): Promise<Counter> {
    return fetch('/api/counter/increment',
        { method: 'POST' })
        .then(res => res.json())
        .then(res => {
            return res as Counter;
        })
}
