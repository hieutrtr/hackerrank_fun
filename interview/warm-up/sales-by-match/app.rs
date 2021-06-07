use std::io::stdin;

fn main() {
    let mut n = String::new();
    let mut socks = String::new();

    stdin().read_line(&mut n).expect("not a string");
    stdin().read_line(&mut socks).expect("not a string");

    let _isocks: &str = &socks[..];
    let _n: i32 = n.trim().parse().unwrap();
    let _socks: Vec<i32> = _isocks.trim().split(" ").into_iter().map(|x| x.parse::<i32>().unwrap()).collect();

    println!("_n :{}", _n);
    println!("_socks : {:?}", _socks);
    let mut skips = Vec::new();
    let mut count = 0;
    for (i, s1) in _socks.iter().enumerate() {
        if !skips.contains(&i) {
            for (j, s2) in _socks.iter().enumerate().skip(i+1) {
                if s1 == s2 {
                    count+=1;
                    skips.push(j);
                    break;
                }
            }
        }
    }
    println!("count: {}", count);
}