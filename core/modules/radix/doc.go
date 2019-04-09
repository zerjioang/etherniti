// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package radix

/*
Provides the radix package that implements a radix tree. The package only provides a single Tree implementation, optimized for sparse nodes.

As a radix tree, it provides the following:

    O(k) operations. In many cases, this can be faster than a hash table since the hash function is an O(k) operation, and hash tables have very poor cache locality.
    Minimum / Maximum value lookups
    Ordered iteration
*/
