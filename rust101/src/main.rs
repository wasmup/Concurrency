fn main() {
    println!("{}", co_sum(1, 8_000_000_000, 8)); // 32000000004000000000
}

fn co_sum(a: u64, b: u64, n: u64) -> u128 {
    let mut threads = Vec::new();
    for (start, end) in cpu_tasks(a, b, n) {
        let thread = std::thread::spawn(move || sum(start, end));
        threads.push(thread);
    }

    let mut total = 0;
    for thread in threads {
        let sum = thread.join().unwrap();
        total += sum;
    }

    total
}

fn sum(a: u64, b: u64) -> u128 {
    let mut sum = 0;
    for i in a..=b {
        sum += i as u128;
    }
    sum
}

fn cpu_tasks(mut a: u64, b: u64, mut n: u64) -> impl Iterator<Item = (u64, u64)> {
    let total_tasks = b - a + 1;
    let tasks_per_core = total_tasks / n;
    let mut remaining_tasks = total_tasks % n;
    if tasks_per_core == 0 {
        n = remaining_tasks;
    }

    (0..n).map(move |_| {
        let mut end = a + tasks_per_core - 1;
        if remaining_tasks > 0 {
            remaining_tasks -= 1;
            end += 1;
        }
        let result = (a, end);
        a = end + 1;
        result
    })
}
