from fpdf import FPDF

def image_to_pdf(input_file, output_file):
    Pdf = FPDF()
    Pdf.add_page()
    Pdf.image(input_file)
    Pdf.output(output_file, "F")

if __name__ == '__main__':
    image_to_pdf("input.png","output.pdf")