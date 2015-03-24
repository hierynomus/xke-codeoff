import scala.io.Source
import java.io.{File, FileWriter}
import scala.math._

object Skeleton {
  import googlecodejam._

  def solveProblem(reader: Iterator[String]) = {
    print(".")
    // TODO implement
    ""
  }

  def main(args: Array[String]) {
    val input: String = args(0)
    val iterator: Iterator[String] = Source.fromFile(input).getLines()
    val file = new File(input.substring(0, input.lastIndexOf(".")) + ".out")
    val nrProblems = iterator.nextInt
    val results = (1 to nrProblems).map(p => "Case #" + p + ": " + solveProblem(iterator))
    val fw = new FileWriter(file)
    try {
      fw.write(results.mkString("\n"))
    } finally {
      fw.close()
    }
  }

  object googlecodejam {
    implicit class RichIterator(val iterator: Iterator[String]) extends AnyVal {
      def trimmedLine = iterator.next().trim()
      def nextInt: Int = trimmedLine.toInt
      def nextIntArray: Array[Int] = nextStringArray map (_.toInt)
      def nextCharArray: Array[Char] = iterator.next().toCharArray
      def nextLongArray: Array[Long] = nextStringArray map (_.toLong)
      def nextDoubleArray: Array[Double] = nextStringArray map (_.toDouble)
      def nextBigDecimalArray: Array[BigDecimal] = nextStringArray map (BigDecimal(_))
      def nextBigIntArray: Array[BigInt] = nextStringArray map (BigInt(_))
      def nextStringArray: Array[String] = trimmedLine split " "
      def nextIntGrid(y: Int): Array[Array[Int]] = (for (i <- 0 until y) yield nextIntArray).toArray
      def nextCharGrid(y: Int): Array[Array[Char]] = (for (i <- 0 until y) yield nextCharArray).toArray

      def skipLines(nr: Int) {
        for (i <- 1 to nr) {
          iterator.next()
        }
      }
    }
  }
}
