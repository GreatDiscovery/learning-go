package util

import (
	"fmt"
	"time"
)

// DefaultFormat allows the Stopwatch.String() function to be
// configured differently to time.Duration if needed.  This is done
// at the global package level to avoid having to do on each Stopwatch
// instance.
var DefaultFormat = func(t time.Duration) string { return t.String() }

// Stopwatch is a structure to hold information about a stopwatch
type Stopwatch struct {
	format  func(time.Duration) string
	elapsed time.Duration
	refTime time.Time
}

// Start returns a pointer to a new Stopwatch struct and indicates
// that the stopwatch has started.
func Start(f func(time.Duration) string) *Stopwatch {
	s := New(f)
	s.Start()

	return s
}

// New returns a pointer to a new Stopwatch struct.
func New(f func(time.Duration) string) *Stopwatch {
	s := new(Stopwatch)
	s.format = f

	return s
}

// Start records that we are now running. If called previously this
// is a no-op.
func (s *Stopwatch) Start() {
	if s.IsRunning() {
		fmt.Printf("WARNING: Stopwatch.Start() IsRunning is true\n")
	} else {
		s.refTime = time.Now()
	}
}

// Stop collects the elapsed time if running and remembers we are
// not running.
func (s *Stopwatch) Stop() {
	if s.IsRunning() {
		s.elapsed += time.Since(s.refTime)
		s.refTime = time.Time{}
	} else {
		fmt.Printf("WARNING: Stopwatch.Stop() IsRunning is false\n")
	}
}

// Reset resets the counters.
func (s *Stopwatch) Reset() {
	if s.IsRunning() {
		fmt.Printf("WARNING: Stopwatch.Reset() IsRunning is true\n")
	}
	s.refTime = time.Time{}
	s.elapsed = 0
}

// String gives the string representation of the duration.
func (s *Stopwatch) String() string {
	// display using local formatting if possible
	if s.format != nil {
		return s.format(s.elapsed)
	}
	// display using package DefaultFormat
	return DefaultFormat(s.elapsed)
}

// SetStringFormat allows the String() function to be configured
// differently to time.Duration for the specific Stopwatch.
func (s *Stopwatch) SetStringFormat(f func(time.Duration) string) {
	s.format = f
}

// IsRunning is a helper function to indicate if in theory the
// stopwatch is working.
func (s *Stopwatch) IsRunning() bool {
	return !s.refTime.IsZero()
}

// Elapsed returns the elapsed time since starting (in time.Duration).
func (s *Stopwatch) Elapsed() time.Duration {
	if s.IsRunning() {
		return time.Since(s.refTime)
	}
	return s.elapsed
}

// ElapsedSeconds is a helper function returns the number of seconds
// since starting.
func (s *Stopwatch) ElapsedSeconds() float64 {
	return s.Elapsed().Seconds()
}

// ElapsedMilliSeconds is a helper function returns the number of
// milliseconds since starting.
func (s *Stopwatch) ElapsedMilliSeconds() float64 {
	return float64(s.Elapsed() / time.Millisecond)
}
