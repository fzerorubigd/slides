def test_try_finally_lock(self):
        counter = [0]
        lock = Lock()
        def loop100(counter):
            for i in xrange(100):
                lock.acquire()
                try:
                    counter[0] += 1
                finally:
                    lock.release()
                time.sleep(0.0001)

