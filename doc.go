/*

Package nrcontext allows you to use the New Relic go-agent with context.Context

Description

The go-agent from New Relic currently (as of 2017-08-30) does not have support
for storing newrelic.Application and newrelic.Transaction in context.Context.

This package is an experiment in doing just that.

Installation

Just go get the package:

    go get -u github.com/TV4/nrcontext

Or more likely: use dep to vendor nrcontext

Links

For more information about New Relic and their go-agent package:

 https://newrelic.com/golang
 https://github.com/newrelic/go-agent

The nrcontext package is based on ideas put forth by Sam Vilain in his article:

	https://medium.com/@gosamv/using-gos-context-library-for-performance-monitoring-aaf25dacb0fe

An alternative to this package is newrelic-context by Maxim Sukharev.

	https://github.com/smacker/newrelic-context

*/
package nrcontext
