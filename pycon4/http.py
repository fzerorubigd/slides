from __go__.net.http import Dir, FileServer, Handle, ListenAndServe

handler = FileServer(Dir("static"))

ListenAndServe('127.0.0.1:8080', handler)
