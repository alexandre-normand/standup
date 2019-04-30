using Go = import "/go.capnp";
@0xb7eae732787dc175;
$Go.package("standup");
$Go.import("github.com/alexandre-normand/standup");

struct Status {
   userID @0 :Text;
   date @1 :Text;
   yesterday @2 :List(Text);
   today @3 :List(Text);
   blockers @4 :List(Text);
}
