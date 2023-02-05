import argparse

parser = argparse.ArgumentParser()
parser.add_argument("-n", "--name", help="Firt name")
parser.add_argument("-l", "--lastname", help="Last name")

args = parser.parse_args()

if args.flag1:
    print("Flag1 is set to: ", args.flag1)
if args.flag2:
    print("Flag2 is set to: ", args.flag2)