CCCOLOR="\033[34m"
LINKCOLOR="\033[34;1m"
SRCCOLOR="\033[33m"
BINCOLOR="\033[37;1m"
MAKECOLOR="\033[32;1m"
ENDCOLOR="\033[0m"


REDIS_SERVER_NAME=main
REDIS_SERVER_OBJ=main.o

all: $(REDIS_SERVER_NAME)

.PHONY: all

# redis-server
$(REDIS_SERVER_NAME): $(REDIS_SERVER_OBJ)
	$(CC) -g -o $@ $^

# Because the jemalloc.h header is generated as a part of the jemalloc build,
# building it should complete before building any other object. Instead of
# depending on a single artifact, build all dependencies first.
%.o: %.c
	$(CC) -c $<

clean:
	rm -rf $(REDIS_SERVER_NAME) *.o *.gcda *.gcno *.gcov redis.info lcov-html Makefile.dep dict-benchmark

.PHONY: clean

