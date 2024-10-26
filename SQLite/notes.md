# SQLite

## lol why tho?
Okay, sqlite is fine to scale with.  If I reach 50k concurrent paying users and need something else at that time, I will be happy $$$

From the docs on [sqlite.org](https://www.sqlite.org/whentouse.html)

> SQLite works great as the database engine for most low to medium traffic websites (which is to say, most websites). The amount of web traffic that SQLite can handle depends on how heavily the website uses its database. Generally speaking, any site that gets fewer than 100K hits/day should work fine with SQLite. The 100K hits/day figure is a conservative estimate, not a hard upper bound. SQLite has been demonstrated to work with 10 times that amount of traffic.

> The SQLite website (https://www.sqlite.org/) uses SQLite itself, of course, and as of this writing (2015) it handles about 400K to 500K HTTP requests per day, about 15-20% of which are dynamic pages touching the database. Dynamic content uses about 200 SQL statements per webpage. This setup runs on a single VM that shares a physical server with 23 others and yet still keeps the load average below 0.1 most of the time.

### Performance Optimizations

#### Write-Ahead Logging (WAL), Disable Synchronous Mode
```
pragma journal_mode = WAL;
pragma synchronous = normal;
pragma journal_size_limit = 6144000;
```






---

## From Chat
SQLite is a highly optimized, self-contained database engine that, while limited in customization compared to other databases, has several settings you can tweak to enhance performance and security. Here are some key configurations you can modify:

### Performance Enhancements

1. **Page Size (`PRAGMA page_size`)**:
   - Default: 4096 bytes (may vary depending on system).
   - Increasing the page size (to, e.g., 8192 or 16384) may improve read/write performance, especially for larger databases with fewer, larger records.

2. **Cache Size (`PRAGMA cache_size`)**:
   - Default: 2000 pages.
   - Increasing cache size can reduce I/O operations by keeping more data in memory, which is helpful for read-heavy applications. Setting it to a negative number specifies memory use in KB (e.g., `PRAGMA cache_size = -2000` for a 2MB cache).

3. **Synchronous Mode (`PRAGMA synchronous`)**:
   - Default: `FULL`.
   - Reducing to `NORMAL` or `OFF` may improve write speed, though this decreases durability (risking data on unexpected crashes). `NORMAL` provides a balance between performance and data integrity.

4. **Journal Mode (`PRAGMA journal_mode`)**:
   - Default: `DELETE`.
   - Setting it to `WAL` (Write-Ahead Logging) allows for concurrent reads and writes, which can improve performance in multi-threaded or multi-user scenarios. Other modes like `MEMORY` and `OFF` can be faster but lack durability.

5. **Locking Mode (`PRAGMA locking_mode`)**:
   - Default: `NORMAL`.
   - `EXCLUSIVE` mode keeps the database locked by the current connection, potentially speeding up writes in single-connection applications, though not recommended for shared environments.

6. **Temp Store (`PRAGMA temp_store`)**:
   - Default: `DEFAULT` (which usually uses the file system).
   - Setting it to `MEMORY` keeps temporary tables and indexes in memory, speeding up operations that use temporary data.

7. **Automatic Indexing (`PRAGMA automatic_index`)**:
   - Default: `ON`.
   - Disabling can improve performance if you carefully design indexes, as automatic indexing can add overhead when working with large data sets or complex queries.

### Security Enhancements

1. **Foreign Key Enforcement (`PRAGMA foreign_keys`)**:
   - Default: `OFF`.
   - Setting this to `ON` enforces foreign key constraints, which is crucial for maintaining data integrity.

2. **Encryption**:
   - **SQLCipher**: This library adds full-database encryption to SQLite, protecting data at rest. Unfortunately, SQLite’s base version doesn’t include encryption by default, but you can compile it with SQLCipher or use the SQLite Encryption Extension (SEE) for secure data handling.

3. **Read-Only Mode**:
   - If you don’t need to write to the database, open it in read-only mode to prevent accidental modifications. This is especially useful for production environments where changes are rare or not permitted.

4. **User Authentication**:
   - SQLite lacks native user authentication, so consider wrapping the database in an API or server that enforces strict authentication and authorization to limit access to the database.

5. **Restrict PRAGMAs in Production**:
   - Some PRAGMAs (e.g., `PRAGMA foreign_keys` and `PRAGMA writable_schema`) can compromise database integrity and should be restricted in production environments.

6. **Sandboxed Environment**:
   - Deploy SQLite in a sandboxed environment where the application and database can’t be accessed by unauthorized users or processes, enhancing data isolation.

Tuning SQLite for your specific workload will depend on factors like data size, expected traffic, read/write frequency, and security requirements. For testing, adjust each PRAGMA incrementally and benchmark the effects to find the best settings.