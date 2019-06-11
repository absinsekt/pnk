use std::{io};

use actix_web::{web, App, Error, HttpRequest, HttpResponse, HttpServer};

use futures::{future::ok, Future};


fn index_async(req: HttpRequest) -> impl Future<Item = HttpResponse, Error = Error> {
    println!("{:?}", req);

    ok(HttpResponse::Ok()
        .content_type("text/html")
        .body(format!("Hi, {}!", "OLOLO!")))
}

fn main() -> io::Result<()> {
    let sys = actix_rt::System::new("basic-example");

    HttpServer::new(|| {
        App::new()
            .service(
                web::resource("/").route(web::get().to_async(index_async)),
            )
    })
    .bind("127.0.0.1:9000")?
    .start();

    sys.run()
}
